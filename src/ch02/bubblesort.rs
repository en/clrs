use std::cmp;
use super::Value;

pub fn bubblesort<T>(a: &mut [Value<T>])
    where T: cmp::PartialOrd + Copy
{
    let n = a.len();
    if n < 2 {
        return;
    }
    for i in 0..n - 1 {
        for j in (i + 1..n).rev() {
            if a[j] < a[j - 1] {
                a.swap(j, j - 1);
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use ::test_cases;

    #[test]
    fn test_bubblesort() {
        macro_rules! test {
            ( $( $i:tt ),* ) => {{
                $(
                    let (mut a, want) = test_cases::SORT.$i;
                    bubblesort(&mut a);
                    assert_eq!(want, a);
                )*
            }};
        }
        test!(0, 1, 2, 3, 4, 5, 6, 7);
    }
}
