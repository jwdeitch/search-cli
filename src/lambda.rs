/// https://iterm2.com/documentation-images.html
/// https://iterm2.com/utilities/imgcat
///
/// ```ignore
/// ESC ] 1337 ; File = [optional arguments] : base-64 encoded file contents ^G
/// ```
use base64::encode;
use get::get_http;
use std::env::var;
use std::io::{Error, ErrorKind, Result, Write};


fn is_iterm() -> bool {
    match var("TERM_PROGRAM") {
        Ok(term) => term == "iTerm.app",
        Err(_) => false,
    }
}

fn print_osc<W: Write>(buf: &mut W) -> Result<()> {
    write!(buf, "{}]", '\u{1B}') // \033
}

fn print_st<W: Write>(buf: &mut W) -> Result<()> {
    write!(buf, "{}", char::from(7)) // \a
}

pub fn inline_image<W>(buf: &mut W, name: &str) -> Result<()>
    where
        W: Write,
{
    if !is_iterm() {
        return Err(Error::new(ErrorKind::Other, "inline images are only supported in iTerm"));
    }

    let image = download_remote_image(String::from(name));

    print_osc(buf)?;
    write!(buf, "1337;File=")?;
    write!(buf, "inline=1")?;

    write!(buf, ":")?;
    write!(buf, "{}", image)?;
    print_st(buf)?;
    write!(buf, "\n")?;
    Ok(())
}

pub fn download_remote_image(path: String) -> String {
    return encode(&get_http(&path).expect("failed retrieving image from WRA API"));
}