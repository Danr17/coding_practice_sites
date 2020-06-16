use std::fmt;

#[derive(Debug, Eq, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let min = (((hours * 60 + minutes) % (24 * 60)) + 24 * 60) % (24 * 60);
        Clock { minutes: min }
    }
    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(0, self.minutes + minutes)
    }
}
impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let hour = self.minutes / 60;
        let minute = self.minutes % 60;
        write!(f, "{:#02}:{:#02}", hour, minute)
    }
}
