$( document ).ready(function() {
  getCategories(renderCategories)
});

function ajaxData(method, path, callback) {
  apiVersion = '/v1'
  $.ajax({
    method: "GET",
    url: apiVersion + path,
    success: callback,
  })
}

function getAjaxData(path, callback) {
  ajaxData('GET', path, callback)
}

function getCategories(callback) {
  getAjaxData('/categories', callback)
}

function renderCategories(categories) {
  var root = $("#sidebar > div > #sidebarCategories")
  root.empty()
  $.each(categories,function(index, category){
    var categoryItem = $('<div />', {
      id: 'category_' + category.ID,
      class: 'bg-dark text-white list-group-item d-flex justify-content-between align-items-center'
    })
    categoryItem.append(`<a categoryID="${category.ID}">${category.Name} </a><a class="badge badge-info"> ${category.UnreadCount || 0}</a>`)
    categoryItem.click(function(){ getItemList(category.ID, renderItemList) })
    root.append(categoryItem)
  })
}

function getItemList(categoryID, callback) {
  categoryID = categoryID || '*'
  getAjaxData(`/categories/${categoryID}/feeds/`, callback)
}

function renderItemList(items) {
  var root = $("#itemList")
  var qHeightDict = {}
  
  root.children(':visible').each(function(){
    childrenHeigth = 0
    console.log($(this))
    $(this).children(':visible').each(function(){
      childrenHeigth += $(this).outerHeight();
    });
    qHeightDict[$(this)[0].id] = childrenHeigth
  })

  console.log(qHeightDict)

}