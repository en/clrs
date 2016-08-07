use std::cmp;

pub fn iterative_linear_search<T: cmp::Eq>(a: &[T], v: &T) -> Result<usize, i32> {
    for i in 0..a.len() {
        if a[i] == *v {
            return Ok(i);
        }
    }
    Err(-1)
}

pub fn recursive_linear_search<T: cmp::Eq>(a: &[T], v: &T, n: usize) -> Result<usize, i32> {
    let n = n - 1;
    if a[n] == *v {
        return Ok(n);
    }
    if n == 0 {
        return Err(-1);
    }
    recursive_linear_search(a, v, n)
}

#[cfg(test)]
mod tests {
    use super::*;
    use ::test_cases;

    #[test]
    fn test_iterative_linear_search() {
        for i in 0..test_cases::SEARCH_RESULT.len() {
            let r = iterative_linear_search(&test_cases::SEARCH, &test_cases::SEARCH_RESULT[i].v);
            assert_eq!(test_cases::SEARCH_RESULT[i].r, r);
        }
    }

    #[test]
    fn test_recursive_linear_search() {
        for i in 0..test_cases::SEARCH_RESULT.len() {
            let r = recursive_linear_search(&test_cases::SEARCH,
                                            &test_cases::SEARCH_RESULT[i].v,
                                            test_cases::SEARCH.len());
            assert_eq!(test_cases::SEARCH_RESULT[i].r, r);
        }
    }
}
