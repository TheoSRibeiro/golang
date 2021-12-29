$('#formulario-cadastro').on('submit', criarUsuario);


function criarUsuario(evento){
    evento.preventDefault();
    console.log("Funcao criarUsuario!");
    
    //Comparar as senhas passadas
    if($('#senha').val() != $('#confirmar-senha').val()){
        Swal.fire("Ops...", "As senhas não coincidem!", "error");
        return;
    }


    //Requisicao AJAX - caso as senhas forem iguais
    $.ajax({
        url: "/usuarios",
        method: "POST",
        data:{
            nome: $('#nome').val(),
            email: $('#email').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val()
        }
    }).done(function(){
        Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "success")
            .then(function(){
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data:{
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(function(){
                    window.location = "/home";
                }).fail(function(){
                    Swal.fire("Ops...", "Erro ao autenticar o Usuário!", "error");
                });
            });

    }).fail(function (erro){
        Swal.fire("Ops...", "Erro ao cadastrar o Usuário!", "error");
    });
}