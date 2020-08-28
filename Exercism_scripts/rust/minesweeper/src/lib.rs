use itertools::iproduct;

pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let height = minefield.len();
    if height == 0 {
        return vec![];
    }
    let width = minefield[0].len();

    minefield.iter().enumerate().map(|(y, line)| {
        let y1 = y.saturating_sub(1);
        let y2 = (y + 1).min(height - 1);

        line.chars().enumerate().map(|(x, c)| {
            let x1 = x.saturating_sub(1);
            let x2 = (x + 1).min(width - 1);

            let nb_mines = iproduct!(x1..=x2, y1..=y2)
                .filter(|(x, y)| minefield[*y].as_bytes()[*x] == b'*')
                .count() as u8;

            match (c, nb_mines) {
                ('*', _) | (_, 0) => c,
                (_, n) => (b'0' + n) as char,
            }
        })
        .collect::<String>()
    })
    .collect::<Vec<_>>()
}