export interface Artist {
    ID: number;
    name: string;
    genre: string;
}

export interface Concert {
    ID: number;
    date: string;
    artist_id: number;
}

export const API_URL_Artist = 'http://localhost:3500'
export const API_URL_Concert = 'http://localhost:3500'