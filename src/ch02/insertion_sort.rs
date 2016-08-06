use std::cmp;

pub fn iterative_insertion_sort<T: cmp::PartialOrd + Copy>(a: &mut [T]) {
    if a.len() < 2 {
        return;
    }
    for j in 1..a.len() {
        let key = a[j];
        // Rust arrays are indexed by usize, which can't be negative.
        // We use a p to record the position where the key will be inserted.
        let mut p: usize = 0;
        for i in (0..j).rev() {
            if a[i] > key {
                a[i + 1] = a[i];
            } else {
                p = i + 1;
                break;
            }
        }
        a[p] = key;
    }
}

pub fn recursive_insertion_sort<T: cmp::PartialOrd + Copy>(a: &mut [T], n: usize) {
    if n < 2 {
        return;
    }
    recursive_insertion_sort(a, n - 1);
    let key = a[n - 1];
    let mut p: usize = 0;
    for i in (0..n - 1).rev() {
        if a[i] > key {
            a[i + 1] = a[i];
        } else {
            p = i + 1;
            break;
        }
    }
    a[p] = key;
}

#[cfg(test)]
mod tests {
    use super::*;
    use ::test_cases;

    #[test]
    fn iterative() {
        macro_rules! iterative {
            ( $( $i:tt ),* ) => {{
                $(
                    let (mut a, want) = test_cases::SORT.$i;
                    iterative_insertion_sort(&mut a);
                    assert_eq!(want, a);
                )*
            }};
        }
        iterative!(0, 1, 2, 3, 4, 5, 6, 7);
    }

    #[test]
    fn recursive() {
        macro_rules! recursive {
            ( $( $i:tt ),* ) => {{
                $(
                    let (mut a, want) = test_cases::SORT.$i;
                    let l = a.len();
                    recursive_insertion_sort(&mut a, l);
                    assert_eq!(want, a);
                )*
            }};
        }
        recursive!(0, 1, 2, 3, 4, 5, 6, 7);
    }
}
