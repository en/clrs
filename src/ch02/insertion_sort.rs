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

    #[test]
    fn iterative() {
        let mut a = [5, 2, 4, 6, 1, 3];
        iterative_insertion_sort(&mut a);
        assert_eq!([1, 2, 3, 4, 5, 6], a);

        let mut a = ['y', 'u', 'a', 'n', 'c', 'h', 'a', 'o'];
        iterative_insertion_sort(&mut a);
        assert_eq!(['a', 'a', 'c', 'h', 'n', 'o', 'u', 'y'], a);

        let mut a: [i32; 0] = [];
        iterative_insertion_sort(&mut a);
        let want: [i32; 0] = [];
        assert_eq!(want, a);

        let mut a = [1];
        iterative_insertion_sort(&mut a);
        assert_eq!([1], a);

        let mut a = [31, 41, 59, 26, 41, 58];
        iterative_insertion_sort(&mut a);
        assert_eq!([26, 31, 41, 41, 58, 59], a);
    }

    #[test]
    fn recursive() {
        let mut a = [5, 2, 4, 6, 1, 3];
        let l = a.len() as i32;
        recursive_insertion_sort(&mut a, l - 1);
        assert_eq!([1, 2, 3, 4, 5, 6], a);

        let mut a = ['y', 'u', 'a', 'n', 'c', 'h', 'a', 'o'];
        let l = a.len() as i32;
        recursive_insertion_sort(&mut a, l - 1);
        assert_eq!(['a', 'a', 'c', 'h', 'n', 'o', 'u', 'y'], a);

        let mut a: [i32; 0] = [];
        let l = a.len() as i32;
        recursive_insertion_sort(&mut a, l - 1);
        let want: [i32; 0] = [];
        assert_eq!(want, a);

        let mut a = [1];
        let l = a.len() as i32;
        recursive_insertion_sort(&mut a, l - 1);
        assert_eq!([1], a);

        let mut a = [31, 41, 59, 26, 41, 58];
        let l = a.len() as i32;
        recursive_insertion_sort(&mut a, l - 1);
        assert_eq!([26, 31, 41, 41, 58, 59], a);
    }
}
