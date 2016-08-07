use std::cmp;
use std::vec::Vec;

#[derive(Debug, Copy, Clone)]
pub enum Value<T> {
    NegInfinity,
    Some(T),
    Infinity,
}

impl<T: cmp::PartialOrd + Copy> PartialEq for Value<T> {
    fn eq(&self, other: &Value<T>) -> bool {
        if let Value::NegInfinity = *self {
            if let Value::NegInfinity = *other {
                return true;
            }
        }
        if let Value::Infinity = *self {
            if let Value::Infinity = *other {
                return true;
            }
        }
        if let Value::Some(x) = *self {
            if let Value::Some(y) = *other {
                return x == y;
            }
        }
        false
    }
}

impl<T: cmp::PartialOrd + Copy> PartialOrd for Value<T> {
    fn partial_cmp(&self, other: &Value<T>) -> Option<cmp::Ordering> {
        match *self {
            Value::NegInfinity => {
                if let Value::NegInfinity = *other {
                    Some(cmp::Ordering::Equal)
                } else {
                    Some(cmp::Ordering::Less)
                }
            }
            Value::Some(x) => {
                match *other {
                    Value::NegInfinity => Some(cmp::Ordering::Greater),
                    Value::Some(y) => x.partial_cmp(&y),
                    Value::Infinity => Some(cmp::Ordering::Less),
                }
            }
            Value::Infinity => {
                if let Value::Infinity = *other {
                    Some(cmp::Ordering::Equal)
                } else {
                    Some(cmp::Ordering::Greater)
                }
            }
        }
    }
}

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
        let q = (p + r) / 2;
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
    fn eq() {
        let a = Value::NegInfinity;
        let b = Value::Some(1);
        let c = Value::Infinity;
        let d = Value::Some(1);
        let e = Value::Some(2);
        assert_eq!(a == a, true);
        assert_eq!(b == b, true);
        assert_eq!(b == d, true);
        assert_eq!(c == c, true);
        assert_eq!(a != b, true);
        assert_eq!(a != c, true);
        assert_eq!(c != b, true);
        assert_eq!(b != e, true);
    }

    #[test]
    fn partial_cmp() {
        let a = Value::NegInfinity;
        let b = Value::Some(1);
        let c = Value::Infinity;
        let d = Value::Some(1);
        let e = Value::Some(2);
        assert_eq!(a < b, true);
        assert_eq!(b < c, true);
        assert_eq!(a < c, true);
        assert_eq!(a <= b, true);
        assert_eq!(b <= c, true);
        assert_eq!(a <= c, true);
        assert_eq!(b < e, true);
        assert_eq!(b <= d, true);
    }

    #[test]
    fn sort() {
        macro_rules! sort {
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
        sort!(0, 1, 2, 3, 4, 5, 6, 7);
    }
}
