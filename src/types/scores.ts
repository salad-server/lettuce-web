export interface Score {
    id: number;
    score: number;
    pp: number;
    acc: number;
    combo: number;
    mods: number;
    c300: number;
    c100: number;
    c50: number;
    geki: number;
    katu: number;
    miss: number;
    rank: string;
    submission_status: number;
    play_date: string;
    perfect_score: boolean;
    map: Beatmap;
}
export interface AdvancedScore {
    id: number;
    uid: number;
    username: string;
    score: number;
    pp: number;
    acc: number;
    combo: number;
    mods: number;
    c300: number;
    c100: number;
    c50: number;
    geki: number;
    katu: number;
    miss: number;
    rank: string;
    submission_status: number;
    play_date: string;
    perfect_score: boolean;
    play_mode: number;
    map: AdvancedBeatmap;

    code?: number;
}
export interface AdvancedBeatmap {
    title: string;
    version: string;
    set_id: number;
    map_id: number;
    artist: string;
    creator: string;
    update: string;
    len: number;
    bpm: number;
    cs: number;
    ar: number;
    od: number;
    hp: number;
    diff: number;
    status: number;
    md5: string;
    max_combo: number;
}

export interface Beatmap {
    map_id: number;
    title: string;
    version: string;
    md5: string;
    set_id: number;
    status: string;
}
