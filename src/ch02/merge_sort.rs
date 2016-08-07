use std::cmp;
use std::vec::Vec;
use super::Value;

pub fn merge<T: cmp::PartialOrd + Copy>(a: &mut [Value<T>], p: usize, q: usize, r: usize) {
    let n1 = q - p + 1;
    let n2 = r - q;
    let mut left: Vec<Value<T>> = Vec::with_capacity(n1 + 1);
    let mut right: Vec<Value<T>> = Vec::with_capacity(n2 + 1);
    left.extend_from_slice(&a[p..p + n1]);
    right.extend_from_slice(&a[q + 1..q + 1 + n2]);
    left.push(Value::Infinity);
    right.push(Value::Infinity);
    let mut i: usize = 0;
    let mut j: usize = 0;
    for k in p..r + 1 {
        if left[i] <= right[j] {
            a[k] = left[i].clone();
            i = i + 1;
        } else {
            a[k] = right[j].clone();
            j = j + 1;
        }
    }
}

pub fn merge_sort<T: cmp::PartialOrd + Copy>(a: &mut [Value<T>], p: usize, r: usize) {
    if p < r {
        let q = (p + r) >> 1;
        merge_sort(a, p, q);
        merge_sort(a, q + 1, r);
        merge(a, p, q, r);
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use ::test_cases;

    #[test]
    fn test_merge_sort() {
        macro_rules! test {
            ( $( $i:tt ),* ) => {{
                $(
                    let (mut a, want) = test_cases::SORT.$i;
                    let l = a.len();
                    let r = if l > 0 { l-1 } else { 0 };
                    merge_sort(&mut a, 0, r);
                    assert_eq!(want, a);
                )*
            }};
        }
        test!(0, 1, 2, 3, 4, 5, 6, 7);
    }
}
