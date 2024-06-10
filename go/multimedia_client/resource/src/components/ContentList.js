import React, { useState } from 'react';
import ContentItem from './ContentItem';

const ContentList = ({ contents, onContentClick }) => {
  const [selectedContentId, setSelectedContentId] = useState(null);

  const handleContentClick = (content) => {
    setSelectedContentId(content.id);
    onContentClick(content);
  };

  return (
    <div>
      <h2>Contents</h2>
      <div className="content-list">
        {contents.map((content, index) => (
          <ContentItem 
            key={content.id} 
            content={content} 
            index={index + 1} 
            onClick={() => handleContentClick(content)} 
            isSelected={content.id === selectedContentId}
          />
        ))}
      </div>
    </div>
  );
};

export default ContentList;
