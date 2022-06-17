export interface Records {
    "ap!std": Record;
    "rx!catch": Record;
    "rx!std": Record;
    "rx!taiko": Record;
    "vn!catch": Record;
    "vn!mania": Record;
    "vn!std": Record;
    "vn!taiko": Record;
}

export interface Record {
    score: Score;
    pp: Score;
}

export interface Score {
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
    play_mode: string;
    map: Map;
}

export interface Map {
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
    status: string;
    md5: string;
    max_combo: number;
}
