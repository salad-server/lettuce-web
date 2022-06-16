export interface Stats {
    id: number;
    total_score: number;
    ranked_score: number;
    pp: number;
    playcount: number;
    acc: number;
    max_combo: number;
    total_hits: number;
    watched_replays: number;
    count_xh: number;
    count_x: number;
    count_sh: number;
    count_s: number;
    count_a: number;
}

export interface Info {
    id: number;
    username: string;
    username_safe: string;
    country: string;
    silence: number;
    register_time: number;
    last_online: number;
    fav_mode: number;
    playstyle: number;
    badge_name: string;
    badge_icon: string;
    bio: string;

    code?: number;
}
