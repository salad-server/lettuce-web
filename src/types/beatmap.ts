// advanced version of Beatmap
export interface Beatmap {
    id: number;
    set_id: number;
    status: string;
    md5: string;
    artist: string;
    title: string;
    version: string;
    creator: string;
    last_update: string;
    total_length: number;
    max_combo: number;
    plays: number;
    passes: number;
    mode: string;
    bpm: number;
    cs: number;
    ar: number;
    od: number;
    hp: number;
    diff: number;
}

// slimmer version of Score
export interface BeatmapScore {
    id: number;
    score: number;
    acc: number;
    max_combo: number;
    mods: number;
    c300: number;
    c100: number;
    c50: number;
    play_date: string;
    miss: number;
    pp: number;
    user: User;
}

export interface User {
    id: string;
    username: string;
    country: string;
}

export interface Fav {
    set_id: number;
    artist: string;
    title: string;
    creator: string;
    children: Child[];
}

export interface Child {
    id: number;
    status: string;
    md5: string;
    version: string;
    mode: string;
    diff: number;
}
