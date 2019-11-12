$( document ).ready(function() {
  $.ajax({
    method: "GET",
    url: "/v1/categories",
    success: renderCategories,
  })
});

function renderCategories(data){
  var root = $("#sidebar > div > div")
  root.empty()
  $.each(data,function(index,childData){
    root.append(`
      <div class="bg-dark text-white list-group-item d-flex justify-content-between align-items-center\">
        <a>Github Trending </a><a class="badge badge-info">16</a>
      </div>
    `)
  })
  return data
}