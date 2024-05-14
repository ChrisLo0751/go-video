import React, { useState } from 'react';
import ReactPlayer from 'react-player';
import { UploadedFile } from './dropzone';
import "./videoCard.css";

interface VideoCardProps {
  video: UploadedFile;
  onDelete: (video: UploadedFile) => void;
  onPlay: (video: UploadedFile) => void;
}

const VideoCard: React.FC<VideoCardProps> = ({ video, onDelete, onPlay }) => {
  const [showPopup, setShowPopup] = useState(false);

  const handlePlayClick = () => {
    setShowPopup(true);
    onPlay(video);
  };

  const handleClosePopup = () => {
    setShowPopup(false);
  };

  return (
    <>
      <div className="video-card" onClick={handlePlayClick}>
        <div className="video-thumbnail">
          <ReactPlayer url={video.preview} width="100%" height="100%" />
        </div>
        <p className='video-name'>视频名称：{video.name}</p>
        <div>
            <button className='play-button' onClick={(e) => { handlePlayClick(); }}>播放</button>
            <button className='delete-button' onClick={(e) => { e.stopPropagation(); onDelete(video); }}>删除</button>
        </div>
      </div>
      {showPopup && (
        <div className="video-popup" onClick={handleClosePopup}>
          <div className="video-popup-content">
            <ReactPlayer url={video.preview} width="100%" height="100%" controls />
            <p className='video-name'>视频名称：{video.name}</p>
          </div>
        </div>
      )}
    </>
  );
};

export default VideoCard;
