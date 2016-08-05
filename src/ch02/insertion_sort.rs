use std::cmp;

pub fn iterative_insertion_sort<T: cmp::PartialOrd + Copy>(a: &mut [T]) {
    if a.len() < 2 {
        return;
    }
    for j in 1..a.len() {
        let key = a[j];
        let mut i: i32 = (j - 1) as i32;
        while i >= 0 && a[i as usize] > key {
            a[(i + 1) as usize] = a[i as usize];
            i = i - 1;
        }
        a[(i + 1) as usize] = key;
    }
}

pub fn recursive_insertion_sort<T: cmp::PartialOrd + Copy>(a: &mut [T], n: i32) {
    if n > 0 {
        recursive_insertion_sort(a, n - 1);
        let key = a[n as usize];
        let mut i: i32 = (n - 1) as i32;
        while i >= 0 && a[i as usize] > key {
            a[(i + 1) as usize] = a[i as usize];
            i = i - 1;
        }
        a[(i + 1) as usize] = key;
    }
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
                    let l = a.len() as i32;
                    recursive_insertion_sort(&mut a, l - 1);
                    assert_eq!(want, a);
                )*
            }};
        }
        recursive!(0, 1, 2, 3, 4, 5, 6, 7);
    }
}
