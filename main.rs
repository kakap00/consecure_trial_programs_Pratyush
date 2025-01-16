use std::io;
use rand::Rng;
use std::cmp::Ordering;

fn main() {
    const blahblah: u32 = 60*60*60;
    println!("Guess the number!");
    let secret_number = rand::thread_rng().gen_range(1..=100);
    //println!("{}", secret_number);
    println!("Input your guess");
    loop{
        println!("input guess:");

        let mut guess = String::new();

        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        let guess: u32 = match guess.trim().parse(){
            Ok(num) => num,
            Err(_) => continue,
        };
    
        println!("you guessed: {}", guess);

        match guess.cmp(&secret_number){
            Ordering::Less => println!("too small"),
            Ordering::Greater => println!("too large"),
            Ordering::Equal => {
                println!("good job");
                break;
            }
        }
    }
}
