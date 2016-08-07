use std::cmp;

pub mod insertion_sort;
pub mod linear_search;
pub mod binary_add;
pub mod selection_sort;
pub mod merge_sort;
pub mod binary_search;
pub mod bubblesort;

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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_eq() {
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
    fn test_partial_cmp() {
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
}
