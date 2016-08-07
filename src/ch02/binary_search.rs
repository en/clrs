use std::cmp;

pub fn iterative_binary_search<T: cmp::PartialOrd>(a: &[T],
                                                   v: &T,
                                                   lo: usize,
                                                   hi: usize)
                                                   -> Result<usize, i32> {
    let (mut lo, mut hi) = (lo, hi);
    while lo <= hi {
        let mid = (lo + hi) >> 1;
        if *v > a[mid] {
            lo = mid + 1;
        } else if *v < a[mid] {
            // ugly
            if mid == 0 {
                return Err(-1);
            }
            hi = mid - 1;
        } else {
            return Ok(mid);
        }
    }
    Err(-1)
}

pub fn recursive_binary_search<T: cmp::PartialOrd>(a: &[T],
                                                   v: &T,
                                                   lo: usize,
                                                   hi: usize)
                                                   -> Result<usize, i32> {
    if lo > hi {
        return Err(-1);
    }
    let mid = (lo + hi) >> 1;
    if *v > a[mid] {
        return recursive_binary_search(a, v, mid + 1, hi);
    } else if *v < a[mid] {
        if mid == 0 {
            return Err(-1);
        }
        return recursive_binary_search(a, v, lo, mid - 1);
    }
    Ok(mid)
}

#[cfg(test)]
mod tests {
    use super::*;
    use ::test_cases;

    #[test]
    fn test_iterative_binary_search() {
        for i in 0..test_cases::SEARCH_RESULT.len() {
            let l = test_cases::SEARCH.len();
            let r = iterative_binary_search(&test_cases::SEARCH,
                                            &test_cases::SEARCH_RESULT[i].v,
                                            0,
                                            l - 1);
            assert_eq!(test_cases::SEARCH_RESULT[i].r, r);
        }
    }

    #[test]
    fn test_recursive_binary_search() {
        for i in 0..test_cases::SEARCH_RESULT.len() {
            let l = test_cases::SEARCH.len();
            let r = recursive_binary_search(&test_cases::SEARCH,
                                            &test_cases::SEARCH_RESULT[i].v,
                                            0,
                                            l - 1);
            assert_eq!(test_cases::SEARCH_RESULT[i].r, r);
        }
    }
}
