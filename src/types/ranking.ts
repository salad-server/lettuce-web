export interface Leaderboard {
    pp: number;
    acc: number;
    total_score: number;
    ranked_score: number;
    playcount: number;
    user: User;
}

export interface User {
    id: number;
    username: string;
    country: string;
}
