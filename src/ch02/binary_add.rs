pub fn binary_add(a: &[i8], b: &[i8], c: &mut [i8]) {
    let mut carry = 0;
    for i in 0..a.len() {
        let s = a[i] + b[i] + carry;
        c[i] = s % 2;
        carry = if s < 2 { 0 } else { 1 }
    }
    c[a.len()] = carry;
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn addition() {
        let a = [0, 0, 0, 1, 1, 1, 0];
        let b = [1, 0, 1, 1, 0, 0, 1];
        let mut c: [i8; 8] = [0; 8];
        binary_add(&a, &b, &mut c);
        assert_eq!([1, 0, 1, 0, 0, 0, 0, 1], c);

        let a = [0, 0, 0, 1, 1, 0, 1];
        let b = [1, 0, 1, 1, 0, 0, 1];
        let mut c: [i8; 8] = [0; 8];
        binary_add(&a, &b, &mut c);
        assert_eq!([1, 0, 1, 0, 0, 1, 0, 1], c);
    }
}
