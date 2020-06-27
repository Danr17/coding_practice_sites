use std::fmt;

#[derive(Debug, Eq, PartialEq)]
pub struct Clock {
    minutes: i32,
}

const MINUTES_PER_DAY: i32 = 24 * 60;
const MINUTES_PER_HOUR: i32 = 60;

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let min = (hours * MINUTES_PER_HOUR + minutes).rem_euclid(MINUTES_PER_DAY);
        Clock { minutes: min }
    }
    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.minutes + minutes)
    }
}
impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let hour = self.minutes.div_euclid(MINUTES_PER_HOUR);
        let minute = self.minutes.rem_euclid(MINUTES_PER_HOUR);
        write!(f, "{:#02}:{:#02}", hour, minute)
    }
}
