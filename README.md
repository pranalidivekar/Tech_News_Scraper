# Tech_News_Scraper

## Description
A web scraper project written in Go meant to address tech dummies who aren't up to date with the latest scoop.
The code scrapes various tech news/blog websites like techcrunch.com, wired.com, theverge.com, ycombinator.com to create an aggregate pdf with the top news items for the day along with their one liner summaries and links to actual articles. The task has been automated by a cron so that the pdf gets emailed to the user on a daily frequency.

## Data Sources
YouTube public API

## Project Structure
/project-directory
|-- /examples
|-- /src
|-- /LICENSE
|-- README.md

- `examples`: Directory to store examples of agrregated news snippets
- `src`: Directory for Python scripts and source code.
- `tests`: Directory for implementing the system.


## Testing
- Update the email.go with the appropriate recipents email ID.
- Run the main.go to fetch results.

## Contributing
Please feel free to fork this repository, make changes, and submit pull requests. Any contributions, no matter how small, are greatly appreciated.

## License
This project is licensed under the MIT License - see the `LICENSE` file for details.

## Contact
If you have any questions or feedback, please feel free to contact me at pranalidivekar@gmail.com.
