export interface Breaks {
    id: string;
    name: string;
    slug: string;
}

export interface Break extends Breaks {
    description: string;
    coordinates: string;
    country: string;
    region: string;
    city: string;
    videoUrl: string | null;
    imageUrls: string[] | null;
}