pub fn nth(n: u32) -> u32 {
    match n {
        0 => 2,
        1 => 3,
        _ => {
            let mut i: u32 = 1;
            let mut idx: u32 = 0;
            loop {
                if is_prime_sqrt(i) {
                    idx += 1
                }
                if idx > n + 1 {
                    break;
                }
                i += 1
            }
            i
        }
    }
}

fn is_prime_sqrt(n: u32) -> bool {
    for i in 2..=((n as f32).sqrt().floor() as u32) {
        if n % i == 0 {
            return false;
        }
    }
    true
}
