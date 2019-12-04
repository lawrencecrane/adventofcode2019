fn main() {
    let (from, to) = parse("137683-596253").unwrap();
    println!("Answer to Part 1 {}", bruteforce_part_1(from, to));
}

fn bruteforce_part_1(from: usize, to: usize) -> usize {
    (from..to)
        .filter(is_valid_part_1)
        .count()
}

fn is_valid_part_1(pw: &usize) -> bool {
    let (_, valid, same) = get_six_digits(pw).iter()
        .fold((0, true, false), |(prev, valid, same), x| {
            (*x, *x >= prev && valid, *x == prev || same)
        });

    valid && same
}

fn get_six_digits(x: &usize) -> [usize; 6] {
    [(x / 100000) % 10,
     (x / 10000) % 10,
     (x / 1000) % 10,
     (x / 100) % 10,
     (x / 10) % 10,
      x % 10]
}

fn parse(input: &str) -> Option<(usize, usize)> {
    let mut splitted = input.split('-');

    match (splitted.next(), splitted.next()) {
        (Some(x), Some(y)) => Some((x.parse().unwrap(), y.parse().unwrap())),
        _ => None
    }
}
