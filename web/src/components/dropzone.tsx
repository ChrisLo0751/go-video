import React, { useState } from 'react';
import { useDropzone } from 'react-dropzone';
import VideoCard from './videoCar';
import './dropzone.css';

export interface UploadedFile extends File {
    preview: string;
}

interface UploadedProps {
    handlerUploadedFiles: (files: File[]) => void;
}

const VideoUploader: React.FC<UploadedProps> = (props) => {
    const [videos, setVideos] = useState<UploadedFile[]>([]);
    const [selectedVideo, setSelectedVideo] = useState<UploadedFile | null>(null);

    const onDrop = (acceptedFiles: File[]) => {
        if (videos.length >= 8 || acceptedFiles.length > 8) {
            alert('视频时长不建议过长，最多仅支持8个视频');
            return;
        }
        const newVideos = acceptedFiles.map((file) =>
            Object.assign(file, {
                preview: URL.createObjectURL(file),
            })
        );
        setVideos((prevVideos) => [...prevVideos, ...newVideos]);
        props.handlerUploadedFiles(acceptedFiles);
    };

    const { getRootProps, getInputProps } = useDropzone({
        accept: { video: ['video/*'] },
        onDrop,
        multiple: true,
    });

    const handleVideoClick = (video: UploadedFile) => {
        setSelectedVideo(video);
    };

    const handleDeleteVideo = (video: UploadedFile) => {
        setVideos((prevVideos) => prevVideos.filter((v) => v !== video));
        setSelectedVideo(null);
    };

    const clearVideos = () => {
        setVideos([]);
        setSelectedVideo(null);
    };

    const uploadFiles = () => { 
        if(videos.length === 0){
            alert('请先上传视频');
            return;
        }
        
    }

    return (
        <div>
            <h2>文件上传</h2>
            <div {...getRootProps({ className: 'dropzone' })}>
                <input {...getInputProps()} />
                <p>拖拽视频或点击上传文件</p>
            </div>
            <h2>视频内容区</h2>
            <div className="button-container">
                <button className="clear-button"  onClick={clearVideos}>全部清空</button>
                <button className="compose-button" onClick={uploadFiles} >合成视频</button>
            </div>
            <div className="video-container">
                {videos.map((video, index) => (
                    <VideoCard video={video} onPlay={handleVideoClick} onDelete={handleDeleteVideo}/>
                ))}
            </div>
        </div>
    );
};

export default VideoUploader;