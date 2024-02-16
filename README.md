# Dating App Documentation
Welcome to the Dating App! This application is testing app that designed to connect individuals in a fun and interactive way, allowing users to discover potential matches based on shared interests, preferences, and geographical proximity. With a user-friendly interface and a range of features, the Dating App provides a platform for users to engage, build connections, and find meaningful relationships.

# Domain-Driven Design (DDD) Structure
This application is developed following the principles of Domain-Driven Design (DDD). DDD promotes a modular and structured approach to building software by organizing code around the business domain.

# Show Of Structure
dating-app
----integration_test
----migrations
----src
--------app
------------dto
------------usecase
--------domain
------------entities
------------repositories
------------value_object
----------------user
--------infra
------------auth
----------------jwt
------------constants
------------helpers
------------models
------------persistence
----------------postgresql
--------interface
------------rest
----------------middleware
----------------response
----------------v1
--------------------mobile_app
------------------------handlers
------------------------requests
------------------------routes
------------------------transformers

