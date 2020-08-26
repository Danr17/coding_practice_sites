#[derive(Debug, PartialEq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    match (_first_list, _second_list) {
        (&[], &[]) => Comparison::Equal,
        (&[], _) => Comparison::Sublist,
        (_, &[]) => Comparison::Superlist,
        (&[..], &[..]) if _first_list == _second_list => Comparison::Equal,
        (&[..], &[..])
            if _first_list.len() > _second_list.len() && contains(_first_list, _second_list) =>
        {
            Comparison::Superlist
        }
        (&[..], &[..])
            if _first_list.len() < _second_list.len() && contains(_second_list, _first_list) =>
        {
            Comparison::Sublist
        }
        (&[..], &[..]) => Comparison::Unequal,
    }
}

fn contains<T>(list1: &[T], list2: &[T]) -> bool
where
    T: PartialEq,
{
    for s in list1.windows(list2.len()) {
        if s == list2 {
            return true;
        }
    }
    return false;
}
