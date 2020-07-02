pub fn build_proverb(list: &[&str]) -> String {
    if list.is_empty() {
        return "".to_string();
    }
    let res_vec: Vec<String> = (0..list.len() - 1)
        .map(|num| {
            format!(
                "For want of a {} the {} was lost.\n",
                list.get(num).unwrap().to_owned(),
                list.get(num + 1).unwrap().to_owned()
            )
        })
        .collect();

    let mut result = res_vec.join("");
    let end: &str = &format!(
        "And all for the want of a {}.",
        list.get(0).unwrap().to_owned()
    );
    result.push_str(end);
    result
}
