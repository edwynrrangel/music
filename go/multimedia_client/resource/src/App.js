import React, { useState, useEffect } from 'react';
import axios from 'axios';
import ContentList from './components/ContentList';
import Player from './components/Player';
import './index.css';

function App() {
  const [contents, setContents] = useState([]);
  const [currentContent, setCurrentContent] = useState(null);

  useEffect(() => {
    async function fetchContents() {
      const { data } = await axios.get('/api/contents');
      setContents(data.data);
    }
    fetchContents();
  }, []);

  const handleContentSelect = (content) => {
    setCurrentContent(content);
  };

  return (
    <div className="App">
      <h1>Multimedia Player</h1>
      <div className="content-container">
        <ContentList contents={contents} onContentClick={handleContentSelect} />
      </div>
      {currentContent && (
        <div className="player-container">
          <Player content={currentContent} />
        </div>
      )}
    </div>
  );
}

export default App;
