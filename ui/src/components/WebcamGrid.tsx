interface WebcamGridProps {
    imageUrls: string[];
    name: string;
}

export function WebcamGrid({imageUrls, name}: WebcamGridProps) {
    return (
        <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4"
             aria-label={`Webcam images for ${name}`}>
            {imageUrls.map((url) => (
                <a key={url} href={url} rel="noopener noreferrer" target="_blank" aria-hidden={true}>
                    <img src={url} className="w-full h-auto rounded-lg shadow-md" alt=""/>
                </a>
            ))}
        </div>
    );
}
