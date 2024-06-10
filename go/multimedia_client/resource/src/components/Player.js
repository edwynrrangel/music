import React, { useEffect, useState, useRef } from 'react';
import 'react-h5-audio-player/lib/styles.css';
import AudioPlayer from 'react-h5-audio-player';

const cleanUpAudio = (audioPlayerRef) => {
  if (audioPlayerRef.current) {
    audioPlayerRef.current.audio.current.pause();
    audioPlayerRef.current.audio.current.removeAttribute('src');
    audioPlayerRef.current.audio.current.load();
  }
};

const Player = ({ content }) => {
  const [audioUrl, setAudioUrl] = useState(null);
  const [error, setError] = useState(null);
  const audioPlayerRef = useRef(null);

  useEffect(() => {
    const fetchAudioData = async () => {
      setError(null); // Reset the error state
      try {
        const response = await fetch(`/api/contents/${content.id}/stream`);
        if (!response.ok) {
          throw new Error('Content not available');
        }
        const reader = response.body.getReader();
        const chunks = [];
        let receivedLength = 0;

        while (true) {
          const { done, value } = await reader.read();
          if (done) {
            break;
          }
          chunks.push(value);
          receivedLength += value.length;
        }

        const blob = new Blob(chunks);
        const url = URL.createObjectURL(blob);
        setAudioUrl(url);
      } catch (err) {
        setError(err.message);
        cleanUpAudio(audioPlayerRef);
      }
    };

    fetchAudioData();

    // Cleanup function to stop the audio and revoke the object URL when the component unmounts or content changes
    return () => {
      cleanUpAudio(audioPlayerRef);
      if (audioUrl) {
        URL.revokeObjectURL(audioUrl);
        setAudioUrl(null);
      }
    };
  }, [content]);

  return (
    <div className="player-content">
      <div className="player-info">
        <div className="player-details">
          <h2>{content.title}</h2>
          <p>{content.creator}</p>
        </div>
        {error && <div className="player-error">{error}</div>}
      </div>
      <AudioPlayer
        ref={audioPlayerRef}
        src={audioUrl}
      />
    </div>
  );
};

export default Player;
