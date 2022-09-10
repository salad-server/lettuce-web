// From ripple (ripple.moe)
const mstr = [
    "NF",
    "EZ",
    "TD",
    "HD",
    "HR",
    "SD",
    "DT",
    "RX",
    "HT",
    "NC",
    "FL",
    "AU",
    "SO",
    "AP",
    "PF",
    "K4",
    "K5",
    "K6",
    "K7",
    "K8",
    "K9",
    "RN",
    "LM",
    "FI",
    "HD",
    "K1",
    "K3",
    "K2",
];

export function modstr(e: number) {
    const n: string[] = [];
    return (
        512 == (512 & e) && (e &= -65),
        16384 == (16384 & e) && (e &= -33),
        mstr.forEach((t, i) => {
            (e & (1 << i)) > 0 && n.push(t);
        }),
        n
    );
}
