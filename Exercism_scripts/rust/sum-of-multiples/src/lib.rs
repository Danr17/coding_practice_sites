use std::collections::HashSet;

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut multiples = HashSet::new();
    for num in factors {
        let mut s = *num;
        loop {
            if s >= limit || s == 0 {
                break;
            }
            multiples.insert(s);
            s += *num;
        }
    }
    multiples.into_iter().sum()
}
