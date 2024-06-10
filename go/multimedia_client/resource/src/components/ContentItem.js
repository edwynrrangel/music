import React from 'react';

const ContentItem = ({ content, index, onClick, isSelected }) => {
  return (
    <div className={`content-item ${isSelected ? 'selected' : ''}`} onClick={onClick}>
      <div className="content-index">{index}</div>
      <div className="content-info">
        <span className="content-title">{content.title}</span>
        <span className="content-creator">{content.creator}</span>
      </div>
      <div className="content-duration">{content.duration}</div>
    </div>
  );
};

export default ContentItem;
