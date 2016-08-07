use std::cmp;
use std::vec::Vec;
use super::Value;

pub fn count_inversions<T: cmp::PartialOrd + Copy>(a: &mut [Value<T>],
                                                   p: usize,
                                                   r: usize)
                                                   -> usize {
    let mut inversions: usize = 0;
    if p < r {
        let q = (p + r) >> 1;
        inversions = inversions + count_inversions(a, p, q);
        inversions = inversions + count_inversions(a, q + 1, r);
        inversions = inversions + merge_inversions(a, p, q, r);
    }
    inversions
}

pub fn merge_inversions<T: cmp::PartialOrd + Copy>(a: &mut [Value<T>],
                                                   p: usize,
                                                   q: usize,
                                                   r: usize)
                                                   -> usize {
    let n1 = q - p + 1;
    let n2 = r - q;
    let mut left: Vec<Value<T>> = Vec::with_capacity(n1 + 1);
    let mut right: Vec<Value<T>> = Vec::with_capacity(n2 + 1);
    left.extend_from_slice(&a[p..p + n1]);
    right.extend_from_slice(&a[q + 1..q + 1 + n2]);
    left.push(Value::Infinity);
    right.push(Value::Infinity);
    let (mut i, mut j): (usize, usize) = (0, 0);
    let mut inversions: usize = 0;
    let mut counted = false;
    for k in p..r + 1 {
        if !counted && right[j] < left[i] {
            inversions = inversions + n1 - i;
            counted = true;
        }
        if left[i] <= right[j] {
            a[k] = left[i];
            i = i + 1;
        } else {
            a[k] = right[j];
            j = j + 1;
            counted = false
        }
    }
    inversions
}

#[cfg(test)]
mod tests {
    use super::*;
    use super::super::Value;

    #[test]
    fn test_count_inversions() {
        let mut a = [Value::Some(0)];
        let l = a.len();
        let i = count_inversions(&mut a, 0, l - 1);
        assert_eq!(0, i);

        let mut a =
            [Value::Some(2), Value::Some(3), Value::Some(8), Value::Some(6), Value::Some(1)];
        let l = a.len();
        let i = count_inversions(&mut a, 0, l - 1);
        assert_eq!(5, i);

        let mut a = [Value::Some(6),
                     Value::Some(5),
                     Value::Some(4),
                     Value::Some(3),
                     Value::Some(2),
                     Value::Some(1)];
        let l = a.len();
        let i = count_inversions(&mut a, 0, l - 1);
        assert_eq!(15, i);

        let mut a = [Value::Some(1),
                     Value::Some(2),
                     Value::Some(3),
                     Value::Some(4),
                     Value::Some(5),
                     Value::Some(6)];
        let l = a.len();
        let i = count_inversions(&mut a, 0, l - 1);
        assert_eq!(0, i);
    }
}
