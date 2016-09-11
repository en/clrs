pub type Stack<T> = Vec<T>;

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_stack() {
        let mut stack: Stack<i32> = Stack::new();
        assert_eq!(true, stack.is_empty());
        assert_eq!(None, stack.pop());
        stack.push(15);
        assert_eq!(false, stack.is_empty());
        stack.push(6);
        stack.push(2);
        assert_eq!(2, stack.pop().unwrap());
        assert_eq!(6, stack.pop().unwrap());
        stack.push(9);
        assert_eq!(9, stack.pop().unwrap());
        assert_eq!(15, stack.pop().unwrap());
    }
}
