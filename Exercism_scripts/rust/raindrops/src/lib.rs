pub fn raindrops(n: u32) -> String {
    let mut out: String = String::new();
    if n % 3 == 0 {
        out += "Pling";
    }
    if n % 5 == 0 {
        out += "Plang";
    }
    if n % 7 == 0 {
        out += "Plong";
    }
    if out.is_empty() {
        return n.to_string();
    }
    out
}
