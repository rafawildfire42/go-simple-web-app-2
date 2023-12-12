function onDelete(id) {
  let response = confirm("Tem certeza que você deseja deletar este produto?")
  if (response) {
    window.location = "/delete?id=" + id
  }
}