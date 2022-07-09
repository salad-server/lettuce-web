# Routes

You can find the routes in `/src/router`.

    [GET] / - Homepage
    [GET] /login - Login page
    [GET] /leaderboard - Leaderboard

    /beatmaps
        [GET] /:id - Beatmap leaderboard

    /info
        [GET] / - Documentation listing
        [GET] /:id - Markdown document

    /profile
        [GET] / - Profile settings
        [GET] /picture - Profile picture settings
        [GET] /:id - User's profile

    /score
        [GET] / - Server records (pp, score, etc...)
        [GET] /:id - Score information
