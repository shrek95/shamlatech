## High Level : Blockchain and smart contracts 
Implement into a blockchain network :

..1 . Renting Property:
..* To rent a property. To include request of property viewing, deposit transfer and regular monthly payments of rent.
..* To rent out a property. To include the listing of said property, tenant verification, tenant acceptance, receipt of monthly rent.

..2 . Property Acquisitions:
..* To place a property up for sale. To include request of property viewing.
..* To manage verified offers of said property.

----------------------------------------------------------------------------------------------------------------

## Chaincode Details:

### User Registering
..1 Register User:
        data_fields:
            ..* User Type(Owner/tenant)
                 ..* First Name
                 ..* Surname
                 ..* DOB
                 ..* Contact Number
                 ..* Email ID
                 ..* Contact Address

### Property Listing
..2 Listing a  property:
        data_fields:
            ..* Property ID
            ..* Owner User ID
            ..* Property status(available/unavailable)
            ..* Property For(Renting/Acquisation)
            ..* Property address
            ..* Property aesthetic attributes
                    ..* type(flat/duplex/villa/triplex)
                    ..* carpet area
                    ..* state of property
                        ..* furnisher(t/f)
                        ..* unfurnished(t/f)
                    ..* Property age
                    ..* rooms
                    ..* balcony
                    ..* parking(Bike/car)

### Property for based monetary details
..3 Property monetary attributes for renting
        data_fields:
            ..* Property ID        
            ..* Monthly rent
            ..* Deposit
            ..* Monthly maintenance
            ..* Additional Expenses

..4 Property monetary attributes for selling
        data_fields
            ..* Property ID
            ..* Price
            ..* Registery charges
            ..* Maintenance expences
            ..* Additional Expenses

### Renting Contact
..5 Renting Listed Property(contract):
        data_fields:
            ..* PropertyID in view
            ..* Deposit Amount
            ..* Tenant User ID
            ..* Owner User ID
            ..* Date
            ..* Tenure of renting(11 months)
### Acquisation contract
..6 Acquisation of Listed Property(contract):
        data_fields:
            ..* Property ID in view
            ..* Price
            ..* Date
            ..* Old Owner ID
            ..* New Owner ID


#Level 1 Details:
        Features offered by the platform are:
            ..1 Renting out a property
            ..2 Buying a property
            ..3 Preview Rented out property
            ..4 Generate virtual contract as rental aggrement
            ..5 Rent a property


#Level 2 Details:
            Flow works in the same format as frontend, edge(middleware+backend processing), backend(hyperledger network)
    Page Layout:
        index.html
            --> Are you a owner(button) or a tenant(button) or a buyer(button).

            --> owner
                    ..* Check user login status
                            ..* NO -> GET: /signup form
                                ..* POST: submit form(button)
                                ..* Generate USER ID                                
                                ..* make DB entry with USER ID as key
                                ..* Response template with UNIQUE USER ID
                                ..* Redirect to /login
                            ..* YES -> GET: /login form
                                ..* POST: submit form(button)
                                ..* Check for credentials
                                ..* response (success/failure)
                                        ..* success : response /home
                                        ..* failure : invalid credentials (retry)
                    ..* ####(success in login)
                    ..* Generate a session for the user
                    ..* GET /home
                            
                            ..* Enter location
                            ..* Submit(button)location
                                    ..* POST: location
                                        ..* store temp.data in redis
                                    ..* reponse: template2.html
                            (template2)        
                            ..* (Button)check_rented_out_property_by_you
                                    ..* GET: /property/rent/usr:USERID
                                    ..* query ledger using USERID
                                    ..* response all property on rent wih USERID
                            ..* (Button)Check_rented_out_properties_by_others_in_location
                                    ..* GET: /property
                                    ..* query DB using location
                                    ..* response all property for rent in <location>
                            ..* (Button)Check_on_sell_property_by_you
                                    ..* GET: /property/sell/usr:USERID
                                    ..* query ledger using USERID
                                    ..* response all property on sell wih USERID
                            ..* (Button)Check_on_sell_property_by_others_in_location
                                    ..* GET: /property
                                    ..* query DB using location
                                    ..* response all property for sell in <location>
                            ..* (Button)Rent_out_a_property_listing
                                ..* GET form template to rent_out_a_property_form:
                                    (form details based on chaincode details listing a property point 2,3)
                                ..* POST: Form submit(button) for renting property
                                    ..* Generate a propertyID
                                    ..* Make a DB entry with keys as propertyID and User ID
                                    ..* Hit the backend with the data in JSON format
                                    ..* initiate fabric INVOKE CreateProperty(...)
                                    ..* respond success/failure
                                        ..* success: redirect to /home
                                        ..* failure: redirect/error:<ERROR>
                            ..* (Button)Sell_a_property_listing
                                ..* GET form template to sell_a_property_listing_form:
                                    (form details based on chaincode details listing a property point 2,4)
                                ..* POST: Form submit(button) for selling property
                                    ..* Generate a propertyID
                                    ..* Make a DB entry with keys as propertyID and User ID
                                    ..* Hit the backend with the data in JSON format
                                    ..* initiate fabric INVOKE CreateProperty(...)
                                    ..* respond success/failure
                                            ..* success: redirect /home
                                            ..* failure: redirect /error:<ERROR>
            --> Tenant
                    ..* Check user login status
                            ..* NO -> GET: /signup form
                                ..* POST: submit form(button)
                                ..* Generate USER ID                                
                                ..* make DB entry with USER ID as key
                                ..* Response template with UNIQUE USER ID
                                ..* Redirect to /login

                            ..* YES -> GET: /login form
                                ..* POST: submit form(button)
                                ..* Check for credentials
                                ..* response (success/failure)
                                        ..* success : response /home
                                        ..* failure : /error
                    ..* ####(success in login)
                    ..* Generate a session for the user
                    ..* GET /home
                            
                            ..* Enter location
                            ..* Submit(button)location
                                    ..* POST: location
                                        ..* store temp.data in redis
                                    ..* reponse: template3.html
                    (template3.html)
                    ..* GET: /property
                    ..* query location from redis
                    ..* response location based rented properties
                        ..* (button)RENT
####                        /* LOGIC TO BE DECIDED */
            --> Buyer 
                    ..* Check user login status
                            ..* NO -> GET: /signup form
                                ..* POST: submit form(button)
                                ..* Generate USER ID                                
                                ..* make DB entry with USER ID as key
                                ..* Response template with UNIQUE USER ID
                                ..* Redirect to /login

                            ..* YES -> GET: /login form
                                ..* POST: submit form(button)
                                ..* Check for credentials
                                ..* response (success/failure)
                                        ..* success : response /home
                                        ..* failure : invalid credentials (retry)
                    ..* ####(success in login)
                    ..* Generate a session for the user
                    ..* GET /home
                            
                            ..* Enter location
                            ..* Submit(button)location
                                    ..* POST: location
                                        ..* store temp.data in redis
                                    ..* reponse: template3.html
                    (template4.html)
                    ..* GET: /property
                    ..* query location from redis
                    ..* response location based for sell properties  
                        ..* (button)BUY
####                        /* LOGIC TO BE DECIDED */

#Level 3 Details:
                        // Working

