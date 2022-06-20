# Load the package
library(bigrquery)

# Ensure you are logged in to make the Bigquery hit
# Run 'gcloud auth application-default login'

# Store the project ID
# Run 'gcloud info' to determine which project you have currently set at your default.
projectid = "PROJECT_ID"

# Set your query
sql <- "SELECT * FROM `bigquery-public-data.usa_names.usa_1910_current` LIMIT 10"

# Run the query; this returns a bq_table object that you can query further
tb <- bq_project_query(projectid, sql)

# Store the first 10 rows of the data in a tibble
sample <-bq_table_download(tb, n_max = 10)

# Print the 10 rows of data
sample