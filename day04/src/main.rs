fn main() {
    println!("{:?}", parse("137683-596253").unwrap());
}

fn parse(input: &str) -> Option<(usize, usize)> {
    let mut splitted = input.split('-');

    match (splitted.next(), splitted.next()) {
        (Some(x), Some(y)) => Some((x.parse().unwrap(), y.parse().unwrap())),
        _ => None
    }
}
