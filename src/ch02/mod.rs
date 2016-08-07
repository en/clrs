use std::cmp;

pub mod insertion_sort;
pub mod linear_search;
pub mod binary_add;
pub mod selection_sort;
pub mod merge_sort;
pub mod binary_search;
pub mod bubblesort;
pub mod inversions;

#[derive(Debug, Copy, Clone)]
pub enum Value<T> {
    NegInfinity,
    Some(T),
    Infinity,
}

impl<T> PartialEq for Value<T>
    where T: cmp::PartialOrd + Copy
{
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

impl<T> PartialOrd for Value<T>
    where T: cmp::PartialOrd + Copy
{
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
        assert!(a == a);
        assert!(b == b);
        assert!(b == d);
        assert!(c == c);
        assert!(a != b);
        assert!(a != c);
        assert!(c != b);
        assert!(b != e);
    }

    #[test]
    fn test_partial_cmp() {
        let a = Value::NegInfinity;
        let b = Value::Some(1);
        let c = Value::Infinity;
        let d = Value::Some(1);
        let e = Value::Some(2);
        assert!(a < b);
        assert!(b < c);
        assert!(a < c);
        assert!(a <= b);
        assert!(b <= c);
        assert!(a <= c);
        assert!(b < e);
        assert!(b <= d);
    }
}
