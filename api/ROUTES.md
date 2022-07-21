# Routes

You can find the routes in `/api/cmd/routes.go`.

    [GET] / - API information (development mode only)
    [POST] /login - Login
    [GET] /leaderboard - Leaderboard

    /@me
        [GET] / - Who is this? (development mode only)
        [POST] /pinned - Pin a score
        [DELETE] /pinned - Unpin a score
        [GET] /profile - Get profile settings
        [POST] /profile - Update profile settings
        [POST] /pfp - Set profile picture
        [DELETE] /pfp - Delete profile picture
        [GET] /fav - Is this beatmap favourited?
        [POST] /fav - Favourite beatmap
        [DELETE] /fav - Unfavourite beatmap

    /docs
        [GET] / - Documentation listing
        [GET] /:id - Markdown document

    /users
        [GET] / - User search
        [GET] /:id - User info
        [GET] /:id/avatar - Profile picture
        [GET] /:id/scores - User scores
        [GET] /:id/stats - User stats
        [GET] /:id/pinned - User pinned
        [GET] /:id/favourite - User favourite

    /score
        [GET] / - Server records (pp, score, etc...)
        [GET] /:id - Score information

    /beatmap/:id
        [GET] / - Beatmap information
        [GET] /sets - Beatmap set information
        [GET] /leaderboard - Beatmap leaderboard
