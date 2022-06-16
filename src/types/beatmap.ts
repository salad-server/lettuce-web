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
