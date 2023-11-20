## RORM An orm style redis key manager
V 0.10
#### This package use [github.com/redis/go-redis/v9](https://github.com/redis/go-redis/v9) as redis client.
#### Reference to [gorm](https://gorm.io/gorm)'s design specifications
### Usage<br/>
 Add Your Model ( just like how to use gorm )<br/><br/>
<code>client := redis.NewClient(&Option) // create your redis client<br/>
&nbsp;rdb := rorm.Open(client, nil) &nbsp; // init<br/>
&nbsp;rdb.AppendModel(User)
</code>
<br/><br/>

Set record (whether it exists or not)<br/><br/>
<code>rdb.Set(&User{Name:"Jack", Mail:"jack@gmail.com", LoginCode:"code-113121"})<br/>
&nbsp;// record has been added to redis server
</code>
<br/><br/>

Get record <br/><br/>
<code>query = &User{Name:"Jack"} // create query <br/>
&nbsp;rdb.Get(&query) // get record<br/>
&nbsp;fmt.Printf(query) // { Jack jack@gmai.com code-113121}
</code>

Delete record IncrBy ... see examples

#### For more detail and usage see [examples](./examples/example)

## Future
1. add sync mode to keep sync with sql databases
2. add more tags and configs to support more functions
3. add more redis actions
