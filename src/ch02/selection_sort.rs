use std::cmp;
use super::Value;

pub fn selection_sort<T>(a: &mut [Value<T>])
    where T: cmp::PartialOrd + Copy
{
    let n = a.len();
    if n < 2 {
        return;
    }
    for j in 0..n - 1 {
        let mut smallest = j;
        for i in j + 1..n {
            if a[i] < a[smallest] {
                smallest = i;
            }
        }
        a.swap(j, smallest);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use ::test_cases;

    #[test]
    fn test_selection_sort() {
        macro_rules! test {
            ( $( $i:tt ),* ) => {{
                $(
                    let (mut a, want) = test_cases::SORT.$i;
                    selection_sort(&mut a);
                    assert_eq!(want, a);
                )*
            }};
        }
        test!(0, 1, 2, 3, 4, 5, 6, 7);
    }
}
