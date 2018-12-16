const tableElem = document.querySelector('.js-users-table');
const tableBodyElem = document.querySelector('.js-table-body');
const formElem = document.querySelector('.js-user-add-form');

const renderTableData = tableRowElem => data => {
  const tableDataElem = document.createElement('td');
  tableDataElem.textContent = data;
  tableRowElem.appendChild(tableDataElem);
}

const renderTableAction = tableRowElem => userId => {
  const renderHtml = `<td><button class="js-delete-user-button" data-id="${userId}">Delete</button></td>`;
  tableRowElem.insertAdjacentHTML('beforeend', renderHtml);
}

async function renderUsers() {
  tableBodyElem.innerHTML = '';
  const users = await getUsers();
  const fragment = document.createDocumentFragment();
  users.forEach(user => {
    const tableRowElem = document.createElement('tr');
    const tableDataRender = renderTableData(tableRowElem);
    const tableActionRender = renderTableAction(tableRowElem);

    tableDataRender(user.user_id);
    tableDataRender(user.user_name);
    tableDataRender(user.email);
    tableActionRender(user.user_id);

    fragment.appendChild(tableRowElem);
  })
  tableBodyElem.appendChild(fragment);
}

async function getUsers() {
  const res = await (await fetch('/users/list')).json();
  return res.data;
}

function handleTableAction(event) {
  const { target } = event;
  if(target.classList.contains('js-delete-user-button')) {
    handleDelete(event);
  }
}

async function handleDelete(event) {
  const { target } = event;
  const userId = target.dataset.id;
  await fetch(`/users?user_id=${userId}`, { method: 'DELETE' });
  renderUsers();
}

async function handleFormSubmit(event) {
  event.preventDefault();
  const formData = new FormData(formElem);
  const user = {}
  for(let item of formData.entries()) {
    user[item[0]] = item[1];
  }
  await fetch('/users', {
    method: 'POST',
    body: JSON.stringify(user)
  });
  renderUsers();
}

tableElem.addEventListener('click', handleTableAction)
formElem.addEventListener('submit', handleFormSubmit)
