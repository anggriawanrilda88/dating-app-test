# Dating App Documentation
Welcome to the Dating App! This application is testing app that designed to connect individuals in a fun and interactive way, allowing users to discover potential matches based on shared interests, preferences, and geographical proximity. With a user-friendly interface and a range of features, the Dating App provides a platform for users to engage, build connections, and find meaningful relationships.

# Domain-Driven Design (DDD) Structure
This application is developed following the principles of Domain-Driven Design (DDD). DDD promotes a modular and structured approach to building software by organizing code around the business domain.

# Show Of Structure

dating-app <br />
----integration_test <br />
----migrations <br />
----src <br />
--------app <br />
------------dto <br />
------------usecase <br />
--------domain <br />
------------entities <br />
------------repositories <br />
------------value_object <br />
----------------user <br />
--------infra <br />
------------auth <br />
----------------jwt <br />
------------constants <br />
------------helpers <br />
------------models <br />
------------persistence <br />
----------------postgresql <br />
--------interface <br />
------------rest <br />
----------------middleware <br />
----------------response <br />
----------------v1 <br />
--------------------mobile_app <br />
------------------------handlers <br />
------------------------requests <br />
------------------------routes <br />
------------------------transformers <br />

