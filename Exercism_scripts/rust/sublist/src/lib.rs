#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    match (_first_list, _second_list) {
        (&[], &[]) => {
            Comparison::Equal
        }
        (&[], _) => {
            Comparison::Sublist
        }
        (_, &[]) => {
            Comparison::Superlist
        }
        (&[..], &[..]) if _first_list ==_second_list => {
            Comparison::Equal
        }
        (&[..], &[..]) if _first_list.contains_slice(_second_list) => {
            Comparison::Superlist
        }
        (&[..], &[..]) if _second_list.contains_slice(_first_list) => {
            Comparison::Sublist
        }
        (&[..], &[..]) => {
            Comparison::Unequal
        }
    }
}

trait Container<T>{
    fn contains_slice(&self, items: &[T]) -> bool;
}

impl<T> Container<T> for &[T]
where T: PartialEq
{
    fn contains_slice(&self, needle: &[T]) -> bool {
        let mut index = 0;
        while index < self.len() {
            
            if self[index..].starts_with(needle) {
                return true
            }
            index += 1;
        }
        return false
    }

}