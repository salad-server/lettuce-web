// https://github.com/osuAkatsuki/bancho.py/blob/master/app/constants/mods.py#L65
// TODO: Deal with key mods & td
const Mods = {
    NOMOD: 0,
    NOFAIL: 1 << 0,
    EASY: 1 << 1,
    TOUCHSCREEN: 1 << 2,
    HIDDEN: 1 << 3,
    HARDROCK: 1 << 4,
    SUDDENDEATH: 1 << 5,
    DOUBLETIME: 1 << 6,
    RELAX: 1 << 7,
    HALFTIME: 1 << 8,
    NIGHTCORE: 1 << 9,
    FLASHLIGHT: 1 << 10,
    AUTOPLAY: 1 << 11,
    SPUNOUT: 1 << 12,
    AUTOPILOT: 1 << 13,
    PERFECT: 1 << 14,
    KEY4: 1 << 15,
    KEY5: 1 << 16,
    KEY6: 1 << 17,
    KEY7: 1 << 18,
    KEY8: 1 << 19,
    FADEIN: 1 << 20,
    RANDOM: 1 << 21,
    CINEMA: 1 << 22,
    TARGET: 1 << 23,
    KEY9: 1 << 24,
    KEYCOOP: 1 << 25,
    KEY1: 1 << 26,
    KEY3: 1 << 27,
    KEY2: 1 << 28,
    SCOREV2: 1 << 29,
    MIRROR: 1 << 30,
};

const KEY_MODS =
    Mods.KEY1 |
    Mods.KEY2 |
    Mods.KEY3 |
    Mods.KEY4 |
    Mods.KEY5 |
    Mods.KEY6 |
    Mods.KEY7 |
    Mods.KEY8 |
    Mods.KEY9;

const MANIA_SPECIFIC_MODS = Mods.MIRROR | Mods.RANDOM | Mods.FADEIN | KEY_MODS;
const OSU_SPECIFIC_MODS = Mods.AUTOPILOT | Mods.SPUNOUT | Mods.TARGET;

export function convertMods(mods: number, vn: number) {
    const DTNC = mods & (Mods.DOUBLETIME | Mods.NIGHTCORE);
    if (DTNC == (Mods.DOUBLETIME | Mods.NIGHTCORE)) {
        mods &= ~Mods.DOUBLETIME;
    } else if (DTNC && mods & Mods.HALFTIME) {
        mods &= ~Mods.HALFTIME;
    }

    if (mods & Mods.EASY && mods & Mods.HARDROCK) {
        mods &= ~Mods.HARDROCK;
    }

    if (mods & (Mods.NOFAIL | Mods.RELAX | Mods.AUTOPILOT)) {
        if (mods & Mods.SUDDENDEATH) {
            mods &= ~Mods.SUDDENDEATH;
        }
        if (mods & Mods.PERFECT) {
            mods &= ~Mods.PERFECT;
        }
    }

    if (mods & (Mods.RELAX | Mods.AUTOPILOT)) {
        if (mods & Mods.NOFAIL) {
            mods &= ~Mods.NOFAIL;
        }
    }

    if (mods & Mods.PERFECT && mods & Mods.SUDDENDEATH) {
        mods &= ~Mods.SUDDENDEATH;
    }

    if (vn != 0) {
        mods &= ~OSU_SPECIFIC_MODS;
    }

    if (vn != 3) {
        mods &= ~MANIA_SPECIFIC_MODS;
    }

    if (vn == 0) {
        if (mods & Mods.AUTOPILOT) {
            if (mods & (Mods.SPUNOUT | Mods.RELAX)) {
                mods &= ~Mods.AUTOPILOT;
            }
        }
    }

    if (vn == 3) {
        mods &= ~Mods.RELAX;
        if (mods & Mods.HIDDEN && mods & Mods.FADEIN) {
            mods &= ~Mods.FADEIN;
        }
    }

    return mods;
}

export function modstr(n: number) {
    const final = [];
    const mods = {
        AP: 8192,
        DT: 64,
        EZ: 2,
        FL: 1024,
        HD: 8,
        HR: 16,
        HT: 256,
        NC: 576,
        NF: 1,
        PF: 16384,
        RX: 128,
        SD: 32,
        SO: 4096,
    };

    type Mod = keyof typeof mods;

    for (const mod of Object.keys(mods)) {
        // gross
        if (mods[mod as Mod] != 0 && n & mods[mod as Mod]) {
            final.push(mod);
        }
    }

    return final;
}
