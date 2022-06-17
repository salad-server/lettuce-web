// more advanced version of AdvancedBeatmap & Beatmap
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

// Even slimmer version of Score
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
