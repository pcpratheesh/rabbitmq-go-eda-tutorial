<html>
    <head>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
        <link rel="stylesheet" href="/public/assets/css/style.css">
        <link rel="stylesheet" href="/public/assets/css/tree.css">
    </head>
    <body>
      <div class="container pt-5 pb-5">
       
      </div>

      


          <div class="container">
             <div class="row justify-content-center align-items-center">
               <div class="col-md-4">
                 <div class="d-block p-2">
                     <div class="card">
                         <div class="card-body">
                            <div class="row">
                              <div class="col">
                                <div class="lead">Producer</div>
                              </div>
                             
                            </div>
                         </div>
                     </div>
                 </div>
               </div>
               
               <div class="col-md-4">
                 <div class="dataflow-b">
                   <div class="data"></div>
                 </div>
               </div>
               <div class="col-md-4">
                  <div id="tree">
                    <div class="branch">
                        <div class="entry"><span>Consumer Group</span>
                        <div class="branch consumers-list">
                          
                        </div>
                      </div>
                    </div>
                  </div>
               </div>
             </div>
           </div>
 
           <div class="container">
             <div class="row g-3 align-items-center">
               <div class="col-auto">
                       <button class="btn btn-primary" id="addNewConsumer">Attach New Consumer</button>
                      <button class="btn btn-secondary" id="run-producer">Distribute Data</button>
               </div>
             </div>
             <div class="row">
               <table class="table">
                 <thead>
                   <tr>
                     <th scope="col">Consumer</th>
                     <th scope="col">Consumed Data</th>
                   </tr>
                 </thead>
                 <tbody id="consumed-data-table">
                  
                 </tbody>
               </table>
             </div>
           </div>
          
    </body>
    <footer>
        <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js" integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo" crossorigin="anonymous"></script>
        <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/js/bootstrap.bundle.min.js" integrity="sha384-MrcW6ZMFYlzcLA8Nl+NtUVF0sA7MsXsP1UyJoMp4YLEuNSfAP+JcXn/tWtIaxVXM" crossorigin="anonymous"></script>
        <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.6.4/jquery.min.js"></script>
        <script src="/public/assets/js/common.js"></script>
        <script src="/public/assets/js/event.js"></script>

    </footer>
</html>