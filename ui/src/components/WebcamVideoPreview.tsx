interface WebcamVideoPreviewProps {
    videoUrl: string;
    name: string;
}

export function WebcamVideoPreview({videoUrl, name}: WebcamVideoPreviewProps) {
    return (
        <div className="mb-8">
            <iframe src={videoUrl} title={`Webcam feed for ${name}`}
                    className="w-full h-full max-w-[800px] aspect-video"
                    referrerPolicy="strict-origin-when-cross-origin" allowFullScreen/>
        </div>
    );
}
