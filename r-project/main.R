# Load the package
library(bigrquery)

# Ensure you are logged in to make the Bigquery hit
# The Bigquery library makes use of your Google Cloud Identity credential flow which provides a
# secure and seamless way to access the relevant data / apis / etc.
# Follow the 3 steps outlined at https://docs.alis.exchange/guides/command-line-interface.html#google-cloud-sdk
# to sort out your authentication.

# Store the project ID
# Run 'gcloud projects list' to determine which project(s) you have available.
projectid <- "YOUR-PROJECT"

# Set your query
sql <- "SELECT * FROM `bigquery-public-data.usa_names.usa_1910_current` LIMIT 10"

# Run the query; this returns a bq_table object that you can query further
tb <- bq_project_query(projectid, sql)

# Store the first 10 rows of the data in a tibble
sample <-bq_table_download(tb, n_max = 10)

# Print the 10 rows of data
sample