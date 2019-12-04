fn main() {
    let (from, to) = parse("137683-596253").unwrap();

    println!("Answer to Part 1 {}", bruteforce_part_1(from, to));
    println!("Answer to Part 2 {}", bruteforce_part_2(from, to));
}

fn bruteforce_part_2(from: usize, to: usize) -> usize {
    (from..to)
        .filter(is_valid_part_2)
        .count()
}

fn bruteforce_part_1(from: usize, to: usize) -> usize {
    (from..to)
        .filter(is_valid_part_1)
        .count()
}

fn is_valid_part_2(pw: &usize) -> bool {
    let (_, valid, sames) = get_six_digits(pw).iter()
        .fold((0, true, [0; 10]), |(prev, valid, mut sames), x| {
            match *x == prev {
                true => {
                    sames[*x] += 1;
                    (*x, valid, sames)
                }
                false => (*x, *x > prev && valid, sames)
            }
        });

    valid && sames.iter().filter(|c| **c == 1).count() > 0
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

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_valid_part2() {
        assert_eq!(is_valid_part_2(&112233), true);
        assert_eq!(is_valid_part_2(&123444), false);
        assert_eq!(is_valid_part_2(&111122), true);
        assert_eq!(is_valid_part_2(&112222), true);
    }
}
